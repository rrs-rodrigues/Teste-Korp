#!/bin/sh

echo "Iniciando instalação do Ansible" 

sudo apt-get update -y && apt-get install ansible -y 

echo "Instalação do Ansible concluída"

sleep 2

echo "Iniciando projeto"

cd ansible

ansible-playbook install-docker.yml

cd ../

echo "Projeto iniciado com sucesso"

