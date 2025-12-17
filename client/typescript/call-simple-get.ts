import { ProductV1ProductServiceApi, Configuration } from '../../gen/ts-client';

const api = new ProductV1ProductServiceApi(new Configuration({ basePath: 'http://localhost:8080' }));

// Fully typed!
const product = await api.productV1ProductServiceGetProduct({ productId: '123' });
console.log(product.name);
