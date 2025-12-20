import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";

// Import generated service definition
import { ProductService } from "../../gen/ts/product_pb";

async function main() {
  // 1. Create the transport
  // This tells Connect how to send the data (HTTP version, base URL, etc.)
  const transport = createConnectTransport({
    baseUrl: "http://localhost:8080", // Pointing to our TS server (or your Go server at 8080)
  });

  // 2. Create the client
  const client = createClient(ProductService, transport);

  // 3. Make the call
  console.log("Calling GetProduct...");
  try {
    const response = await client.getProduct({
      productId: "ts-client-123",
    });

    console.log("Response received:");
    console.log(` - Name: ${response.name}`);
    console.log(` - Price: $${response.price}`);
  } catch (err) {
    console.error("Error calling server:", err);
  }
}

main();
