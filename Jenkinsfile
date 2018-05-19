pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh 'CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o auiHash .'
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
    stage('Integration tests') {
      steps {
        sh '''#/bin/bash

./auiHash &

HTTP_RESPONSE=`curl -i -H "Accept: text/plain" -H "Content-Type: text/plain" -X GET http://localhost:8008/all/String_for_testing`

EXPECTED_RESPONSE="
1f9be8d2262152abbf9c595fe8651ce9
49f0edf87144e8aef8fcf43753cbd7a2497998b2
82022c87bb14169295b5b13688404f013f4c39011c204ae358bff579
77c307a66057925a284f6fe6346b5a89bd11e93be3a39e0da43b37fdf05d61d6
40aa6fb476e83ac0a82aac3484da942a5fa417bdf376f115298cfd28b9d0093cd282fa678d8549f3624108c0a27fb7bb
f9e6aa9514902a0362c64c9849b41bab1525d4d1732e8807de8a380015996eb6ab57e5a613845add6524f4cdd2dc5c9b8ac86343c1977eb8ae2fe150b8697771"

TRIM_RESPONSE=`echo -e $HTTP_RESPONSE | tr \' \' \'\\n\' | tail -6 | xargs -n6`
TRIM_EXPECTED_RESPONSE=`echo -e $EXPECTED_RESPONSE | tr \' \' \'\\n\' | tail -6 | xargs -n6`

echo "**************************************************************************************************************************"
echo "$TRIM_RESPONSE"
echo "**************************************************************************************************************************"
echo "$TRIM_EXPECTED_RESPONSE"
echo "**************************************************************************************************************************"

if [ "$TRIM_RESPONSE" = "$TRIM_EXPECTED_RESPONSE" ]
then
    echo "TEST PASSED!"
    exit 0
else
    echo "*** TEST FAILED!"
    exit 1
fi

kill %1'''
      }
    }
    stage('Clean-up') {
      steps {
        sh 'go test -cover -coverprofile=c.out'
        sh 'go tool cover -html=c.out -o coverage.html'
      }
      publishHTML (target: [
            allowMissing: false,
            alwaysLinkToLastBuild: false,
            keepAll: true,
            reportDir: '.',
            reportFiles: 'coverage.html',
            reportName: "Coverage Report"
          ])
    }
    stage('Clean-up') {
      steps {
        sh 'rm auiHash'
      }
    }
  }
}