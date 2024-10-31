#!/bin/bash

name="$1"
email="$2"
phone="$3"
./cmd/generate_pdf/generate_pdf "$name" "$email" "$phone"
