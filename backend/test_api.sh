#!/bin/bash

# Simple test script for the image generation API

# Test 1: Health check
echo "Testing health endpoint..."
curl -X GET http://localhost:8080/health
echo ""

# Test 2: Get available models
echo "Getting available models..."
curl -X GET http://localhost:8080/api/models
echo ""

# Test 3: Generate image with Gemini
echo "Testing image generation with Gemini..."
curl -X POST http://localhost:8080/api/generate \
  -H "Content-Type: application/json" \
  -d '{
    "prompt": "A nano banana dish in a fancy restaurant with a Gemini theme",
    "model": "gemini-2.0-flash",
    "count": 1,
    "steps": 50,
    "cfg_scale": 7.5
  }'
echo ""

# Test 4: Test with GPT-4 enhanced prompt
echo "Testing image generation with GPT-4 enhancement..."
curl -X POST http://localhost:8080/api/generate \
  -H "Content-Type: application/json" \
  -d '{
    "prompt": "A nano banana dish in a fancy restaurant",
    "model": "gpt-4-image",
    "count": 1
  }'
echo ""
