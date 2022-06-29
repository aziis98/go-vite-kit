import { defineConfig } from 'vite'

export default defineConfig({
    build: {
        rollupOptions: {
            input: {
                '/index': 'index.html',
            },
        },
    },
    server: {
        proxy: {
            '/api': 'http://127.0.0.1:4000/',
        },
    },
})
