package room

// What type of change occurred
type ChangeType string

const (
    PLAYER_JOINED  ChangeType = "joined"  // Player joined
    PLAYER_PLAYING     ChangeType = "playing"     // Player started playing
    PLAYER_DISCONNECTED  ChangeType = "disconnected"  // Player disconnected
)

// Represents a single change from a room
type RoomChange struct {
    Type              ChangeType
    LatestRoom *Room
    //LatestChangeTuple OrderShelfTuple
}

// A change watcher will watch changes from the room
// Changes from the room are buffered in the watcher and stored until consumption by a consumer
// Changes are read in FIFO order
// Changes start buffering once the watcher is created, watcher will not contain changes previous to creation
// Once a change is consumed by a consumer, no other consumer can take the change
// This is not based on channels as we want a buffered slice of changes
type ChangeWatcher struct {
    RoomID string
    in  chan *RoomChange
    out chan *RoomChange
    // Can be used as an identifier
    number int
    // Only written to from routine started by start
    buffer []*RoomChange
}

func (c *ChangeWatcher) SetNumber(number int) {
    c.number = number
}

func (c *ChangeWatcher) Number() int {
    return c.number
}

// Read a single change from the watcher
// Returns the change and whether or not the watcher has been stopped
// Reads in FIFO order
// Removes change from buffer
// Blocks until a change is available
func (c *ChangeWatcher) ReadChange() (*RoomChange, bool) {
    out, ok := <-c.out

    return out, !ok
}

// Add a change to the watcher
// Should be non blocking as long as ChangeWatcher is start'ed
func (c *ChangeWatcher) AddChange(change *RoomChange) {
    c.in <- change
}

// Start the watcher's internal loop
// This loop creates a lock-less infinitely sized channel, non-blocking in channel
// Will be started in the background
func (c *ChangeWatcher) Start() {

    c.in = make(chan *RoomChange)
    c.out = make(chan *RoomChange)
    c.buffer = nil
    go func() {

        for c.in != nil {
            select {
            case in, ok := <-c.in:
                if !ok {
                    c.in = nil
                } else {
                    c.buffer = append(c.buffer, in)
                }

            case c.outC() <- c.curV():
                // Remove element from buffer
                c.buffer = c.buffer[1:]
            }
        }

        close(c.out)
        c.buffer = nil

    }()
}

// Get the value we want to send to the out channel
// Returns the last element in the slice, if not empty
// Otherwise returns nil
func (c *ChangeWatcher) curV() *RoomChange {
    if len(c.buffer) == 0 {
        return nil
    }
    return c.buffer[0]
}

// Use a channel trick to create an infinitely blocking channel when buffer is empty
// Only used in start
func (c *ChangeWatcher) outC() chan *RoomChange {
    if len(c.buffer) <= 0 {
        return nil
    }
    return c.out
}

// Stop our watcher
// All changes in buffer will be discarded
func (c *ChangeWatcher) stop() {
    close(c.in)
}
