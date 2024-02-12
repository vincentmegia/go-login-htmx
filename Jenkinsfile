pipeline {
  agent {
    docker {
      image 'jenkins/jenkins:lts'
    }
  }
  stages {
    stage('Build') {
      steps {
        ./build.sh
      }
    }
  }
}
