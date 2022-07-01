import { basename } from 'path'
import { defineConfig } from 'vite'

// Routes
import routes from './routes.js'

const entryPoints = Object.fromEntries(
    Object.values(routes).map(path => [basename(path, '.html'), path])
)

export default defineConfig({
    build: {
        rollupOptions: {
            input: entryPoints,
        },
    },
    server: {
        proxy: {
            '/api': 'http://127.0.0.1:4000/',
        },
    },
})
