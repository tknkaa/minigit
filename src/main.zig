const std = @import("std");

pub fn main() !void {
    const ns = [4]u8{48, 24, 12, 6};
    const stdout = std.io.getStdOut().writer();
    try stdout.print("{d}\n", .{ ns[2] });
}