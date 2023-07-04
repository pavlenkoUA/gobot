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