pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh 'go build'
      }
    }
    stage('Unit tests') {
      parallel {
        stage('MD5') {
          steps {
            sh 'go test -run MD5'
          }
        }
        stage('SHA1') {
          steps {
            sh 'go test -run SHA1'
          }
        }
        stage('SHA224') {
          steps {
            sh 'go test -run SHA224'
          }
        }
        stage('SHA256') {
          steps {
            sh 'go test -run SHA256'
          }
        }
        stage('SHA384') {
          steps {
            sh 'go test -run SHA384'
          }
        }
        stage('SHA512') {
          steps {
            sh 'go test -run SHA512'
          }
        }
      }
    }
  }
}