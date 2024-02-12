pipeline {
  checkout scmGit(branches: [[name: '*/master']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/vincentmegia/go-login-htmx.git']])
  agent {
    docker {
      image 'jenkins/jenkins:lts'
    }
  }
  stages {
    stage('Checkout sources') {
      echo "hello"
    }
  }
}
