{ pkgs }:
pkgs.mkShell {
  # Add build dependencies
  packages = with pkgs; [
    zig
  ];

  # Add environment variables
  env = { };

  # Load custom bash code
  shellHook = ''
    export ZIG_GLOBAL_CACHE_DIR=$HOME/.cache/zig
  '';
}
