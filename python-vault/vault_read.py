import os
import hvac

client = hvac.Client(
  url="http://localhost:8200",
  token="s.DAqYIayFGSSSoqQwAAHxX9be"
)

data = client.read('demo1/auth')['data']

print(data)

for key, value in data.items():
    print (key, value)