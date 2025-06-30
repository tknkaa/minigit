const std = @import("std");

fn add(x: u8, y: u8) *const u8 {
    const result = x + y;
    return &result;
}

pub fn main() !void {
    const r = add(5, 27);
    std.debug.print("memory address: {}\n", .{r});
    std.debug.print("value: {d}\n", .{r.*});
}
