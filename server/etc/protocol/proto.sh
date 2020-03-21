SRC_DIR="."
DST_DIR="."

protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/bookmark.proto
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/stat.proto
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/tweet.proto
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/user.proto
