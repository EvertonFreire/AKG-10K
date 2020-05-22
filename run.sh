#!/bin/bash

sudo apt install ansible && ansible-galaxy collection install community.kubernetes && ansible-playbook ansible/ansible-playbook.yaml -i localhost, $1


