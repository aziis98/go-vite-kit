import { defineConfig } from 'vite'
import { getBuildRoutes, getDevelopmentRoutes } from './meta/routes.js'

import { join } from 'path'

export default defineConfig(async config => {
    const routes = config.command === 'build' ? await getBuildRoutes() : await getDevelopmentRoutes()
    const entryPoints = Object.values(routes)
    console.log('Found entrypoints:', entryPoints)

    return {
        root: './frontend',
        build: {
            outDir: '../out/frontend',
            emptyOutDir: true,
            rollupOptions: {
                input: entryPoints.map(e => join('./frontend', e)),
            },
        },
        server: {
            proxy: {
                '/api': 'http://127.0.0.1:4000/',
            },
        },
    }
})
