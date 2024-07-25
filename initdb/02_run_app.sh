#!/bin/bash

/app/wait-for-it.sh mysql:3306 -- /app/main &