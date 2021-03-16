resource "helm_release" "golang-api" {
  name       = "golang-api"
  chart      = "./../helm"

}