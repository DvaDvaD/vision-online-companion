{
  description = "Go development environment with custom Go tools";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/5629520edecb69630a3f4d17d3d33fc96c13f6fe"; # Commit with Go 1.23.0
  };

  outputs = { self, nixpkgs }: let
    allSystems = [
      "x86_64-linux"
      "aarch64-linux"
      "x86_64-darwin"
      "aarch64-darwin"
    ];

    forAllSystems = f:
      nixpkgs.lib.genAttrs allSystems (system:
        f {
          pkgs = import nixpkgs { inherit system; };
        });
  in {
    devShells = forAllSystems ({ pkgs }: {
      default = pkgs.mkShell {
        packages = [
          pkgs.go
          pkgs.air
          pkgs.goose
          pkgs.sqlc
          pkgs.oapi-codegen
          pkgs.migrate
          pkgs.k6
        ];

        shellHook = ''
          export GOPRIVATE=github.com/portierglobal
          export PORT=8888

          # List of module directories
          MODULE_DIRS="api business database"

          # Run go mod tidy in each module directory
          for dir in $MODULE_DIRS; do
            if [ -f "$dir/go.mod" ]; then
              echo "Running 'go mod tidy' in $dir"
              (cd "$dir" && go mod tidy)
            else
              echo "No go.mod found in $dir; skipping."
            fi
          done

          chmod +x ./api/scripts/generate-api.sh
          chmod +x ./api/scripts/generate-test.sh

          echo "Development environment ready. Go and Go tools installed."
        '';
      };
    });
  };
}
