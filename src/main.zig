const std = @import("std");

pub fn main() void {
    const level: u8 = 2;
    const category = switch (level) {
        1, 2 => "beginner",
        3 => "professional",
        else => {
            @panic("Not supported");
        },
    };
    std.debug.print("{s}\n", .{category});
}