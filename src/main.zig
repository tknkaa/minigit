const user = @import("user.zig");

pub fn main() !void {
    const u = user.User.init(1, "pedro", "emai@gmail.com");
    u.print_name();
}