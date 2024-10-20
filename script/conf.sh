Target="cfip"
Docker="king011/cfip"
Dir=$(cd "$(dirname $BASH_SOURCE)/.." && pwd)
Version="v0.0.1"
Platforms=(
    darwin/amd64
    windows/amd64
    linux/amd64
)