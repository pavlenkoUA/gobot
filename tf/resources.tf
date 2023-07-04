resource "helm_release" "flux" {
    name    = "flux"
    repository = "https://charts.fluxcd.io"
    chart       = "flux"
    version     = "1.3.0"
    namespace   = "flux-system"

    set {
      name = "git.url"
      value = "git@github.com:pavlenkoua/flux_demo.git"
    }

    set {
      name = "git.path"
      value = "/"
    }

    values = [file("${path.module}/secret-manifest.enc.yaml")]
}
