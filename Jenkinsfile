pipeline {
  node {
    docker {
      image 'jenkins/jenkins:lts'
    }
  }
  stages {
    stage('Checkout sources') {
      echo "hello"
      ./build.sh
    }
  }
}
