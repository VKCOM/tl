using Go = import "/go.capnp";
@0x85d3acc39d94e0f8;
$Go.package("speedcapnp");
$Go.import("speedcapnp");

struct Point {
    x @0 :UInt32;
    y @1 :UInt32;
    z @2 :UInt32;
}

# move file into speedcapnp folder after generation
# user@boot3879:~/go/src/gitlab.corp.mail.ru/vkgo/vkgo/projects/vktl/internal/msgpack$ capnp compile -I /home/user/go/pkg/mod/capnproto.org/go/capnp/v3@v3.0.1-alpha.1/std -ogo speed.capnp
# mv speed.capnp.go speedcapnp/
