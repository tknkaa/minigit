const std = @import("std");

pub fn main() void {
    const a = [5]u8{ 'h', 'e', 'l', 'l', 'o' };
    const b = [_]u8{ 'w', 'o', 'r', 'l', 'd' };
    std.debug.print("{any}", .{a});
    std.debug.print("{any}", .{b});
    const length = b.len;
    std.debug.print("{}", .{length});
}
