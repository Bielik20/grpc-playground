import type { ConnectRouter } from "@connectrpc/connect";
import { connectNodeAdapter } from "@connectrpc/connect-node";
import * as http from "http";
import { ProductService, type GetProductRequest } from "../../gen/ts/product_pb";

// 1. Implement the logic
// The router function allows you to register multiple services
function routes(router: ConnectRouter) {
  router.service(ProductService, {
    async getProduct(req: GetProductRequest, context) {
      console.log(`[TS Server] Received request for ID: ${req.productId} ContentType: ${context.requestHeader.get("content-type")} HttpMethod: ${context.requestMethod}`);

      if (!req.productId) {
        throw new Error("Product ID is required");
      }

      // Return the raw object matching the Product message interface
      return {
        id: req.productId,
        name: "TypeScript Gadget",
        price: 42.5,
      };
    },
  });
}

// 2. Create the Node.js HTTP handler
// This adapter converts standard Node HTTP requests into Connect requests
const handler = connectNodeAdapter({
  routes,
  // If you want to support gRPC-web or standard gRPC, you can configure that here.
  // By default, it supports Connect (JSON/Binary) and gRPC-Web.
});

// 3. Start the server
const port = 8080;
http.createServer(handler).listen(port, () => {
  console.log(`TypeScript Server running on http://localhost:${port}`);
});
