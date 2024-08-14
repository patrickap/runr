PROJECT_NAME := "patrickap/runr"
PROJECT_VERSION := "VERSION"

[private]
get_version:
  @cat {{PROJECT_VERSION}}

[private]
set_version version:
  @echo {{version}} > {{PROJECT_VERSION}}

[private]
backup_version:
  @cp {{PROJECT_VERSION}} {{PROJECT_VERSION}}.bak

[private]
restore_version:
  @cp {{PROJECT_VERSION}}.bak {{PROJECT_VERSION}}

[private]
go-build:
  @go mod download
  @go build -ldflags "-X 'github.com/patrickap/runr/m/v2/cmd.version=$(just get_version)'" -o ./build/runr

[private]
git-publish:
  @git add . -- ':!{{PROJECT_VERSION}}.bak'
  @git commit -m "chore(release): $(just get_version)"
  @git push

[private]
clean_up:
  @rm {{PROJECT_VERSION}}.bak

[private]
release_patch:
  @just backup_version
  @just set_version $(just get_version | awk -F. -v OFS=. '{$3++; print}')
  @just go-build && just git-publish || just restore_version
  @just clean_up

[private]
release_minor:
  @just backup_version
  @just set_version $(just get_version | awk -F. -v OFS=. '{$2++; $3=0; print}')
  @just go-build && just git-publish || just restore_version
  @just clean_up

[private]
release_major:
  @just backup_version
  @just set_version $(just get_version | awk -F. -v OFS=. '{$1++; $2=0; $3=0; print}')
  @just go-build && just git-publish || just restore_version
  @just clean_up

release type:
  @just release_{{type}}
