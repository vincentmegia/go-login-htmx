pipeline {
  agent {
    docker {
      image 'jenkins/jenkins:lts'
      args '--name ${BUILD_TAG}'
    }
  }
  stages {
    stage('Package') {
      steps {
        sh '''-v $(which docker):/usr/bin/docker run \
                --user root \
                --volumes-from "${BUILD_TAG}" \
                buildpacksio/pack:latest build "test-imag"
                  --path "${WORKSPACE}"'''
      }
    }
  }
}
