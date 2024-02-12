pipeline {
  agent {
    docker {
      image 'jenkins/jenkins:lts'
    }
  }
  stages {
    stage('Checkout sources') {
      node {
        checkout scm
      }
    }
  }
}
