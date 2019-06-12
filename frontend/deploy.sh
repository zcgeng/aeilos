# !/bin/sh

npm run build;
# sed -i '' 's+/static/+/aeilos/static/+g' ./build/index.html;
sed -i '' 's+ws://localhost:8000/ws/+wss://changgeng.me/ws/+g' ./build/static/js/main*;

ssh -t hp 'rm -r ~/aeilos/frontend/build-bak || true; mv ~/aeilos/frontend/build ~/aeilos/frontend/build-bak || true;'
scp -r build hp:~/aeilos/frontend/; 
