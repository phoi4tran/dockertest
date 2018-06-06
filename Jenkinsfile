pipeline {
    agent { docker { image 'golang:1.5.0' } }
    stages {
        stage('build') {
            steps {
                sh 'docker image build -t jenkins-docker .'
            }
        }
    }
}
