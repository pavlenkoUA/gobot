module "gke_cluster" {
  source         = "github.com/pavlenkoua/tf-google-gke-cluster"
  GOOGLE_REGION  = var.GOOGLE_REGION
  GOOGLE_PROJECT = var.GOOGLE_PROJECT
  GKE_NUM_NODES  = 1
}
terraform {
  backend "gcs" {
    bucket = "bucket-1test-devops"
    prefix = "terraform/state"
  }
}

module "fluxcd_flux_bootstrap" {
  source            = "github.com/den-vasyliev/tf-fluxcd-flux-bootstrap"
  github_repository = "${var.GHCR_USERNAME}/${var.FLUX_GITHUB_REPO}"
  private_key       = module.tls_private_key.private_key_pem
  github_token      = var.GHCR_TOKEN
}

module "github_repository" {
  source                   = "github.com/pavlenkoua/tf-github-repository"
  github_owner             = var.GHCR_USERNAME
  github_token             = var.GHCR_TOKEN
  repository_name          = var.FLUX_GITHUB_REPO
  public_key_openssh       = module.tls_private_key.public_key_openssh
  public_key_openssh_title = "flux0"
}

module "tls_private_key" {
  source = "github.com/den-vasyliev/tf-hashicorp-tls-keys"
}

module "kind_cluster" {
  source = "github.com/den-vasyliev/tf-kind-cluster"
}