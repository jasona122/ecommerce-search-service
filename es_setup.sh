#!/bin/sh

echo "Running ES Setup..."
apk --no-cache add curl
if curl -I "http://elasticsearch:9200/products" | grep -q "200 OK"; then
  echo "Products index already created, moving to next step..."
else
  curl -X PUT "http://elasticsearch:9200/products" -H "Content-Type: application/json" -d '{
    "mappings": {
      "properties": {
        "name": { "type": "text" },
        "description": { "type": "text" },
        "category": { "type": "keyword" },
        "image_url": { "type": "text" },
        "price": { "type": "float" },
        "quantity": { "type": "integer" },
        "service_area_id": { "type": "keyword" },
        "shop_name": { "type": "text" }
      }
    }
  }'
fi
echo "Setup done!"

