const std = @import("std");

pub fn main() !void {
    var stdout = std.fs.File.stdout();
    _ = try stdout.write("Hello, world!\n");
}
