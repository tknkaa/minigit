{ pkgs }:
pkgs.mkShell {
  # Add build dependencies
  packages = with pkgs; [
    cowsay
  ];

  # Add environment variables
  env = { };

  # Load custom bash code
  shellHook = ''

  '';
}
