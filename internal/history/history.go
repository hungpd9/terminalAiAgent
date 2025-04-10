package history

import (
    "bytes"
    "encoding/gob"
    "time"

    "github.com/boltdb/bolt"
)

type Command struct {
    Text      string
    Timestamp time.Time
}

type History struct {
    db *bolt.DB
}

func NewHistory(dbPath string) (*History, error) {
    db, err := bolt.Open(dbPath, 0600, nil)
    if err != nil {
        return nil, err
    }

    // Tạo bucket nếu chưa tồn tại
    err = db.Update(func(tx *bolt.Tx) error {
        _, err := tx.CreateBucketIfNotExists([]byte("commands"))
        return err
    })
    if err != nil {
        return nil, err
    }

    return &History{db: db}, nil
}

func (h *History) Add(cmd string) error {
    return h.db.Update(func(tx *bolt.Tx) error {
        bucket := tx.Bucket([]byte("commands"))
        id, _ := bucket.NextSequence()
        c := Command{
            Text:      cmd,
            Timestamp: time.Now(),
        }

        var buf bytes.Buffer
        if err := gob.NewEncoder(&buf).Encode(c); err != nil {
            return err
        }

        return bucket.Put(itob(id), buf.Bytes())
    })
}

func (h *History) GetAll() ([]string, error) {
    var commands []string
    err := h.db.View(func(tx *bolt.Tx) error {
        bucket := tx.Bucket([]byte("commands"))
        cursor := bucket.Cursor()

        for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
            var cmd Command
            if err := gob.NewDecoder(bytes.NewReader(v)).Decode(&cmd); err != nil {
                return err
            }
            commands = append(commands, cmd.Text)
        }
        return nil
    })
    if err != nil {
        return nil, err
    }

    // Đảo ngược danh sách để hiển thị lệnh mới nhất trước
    for i, j := 0, len(commands)-1; i < j; i, j = i+1, j-1 {
        commands[i], commands[j] = commands[j], commands[i]
    }
    return commands, nil
}

func (h *History) Close() {
    h.db.Close()
}

// itob converts uint64 to byte slice
func itob(v uint64) []byte {
    buf := make([]byte, 8)
    for i := 0; i < 8; i++ {
        buf[i] = byte(v >> (i * 8))
    }
    return buf
}