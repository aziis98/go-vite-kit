import { defineConfig } from 'vite'
import { getBuildRoutes } from './meta/routes.js'

import { join } from 'path'

export default defineConfig(async () => {
    const routes = await getBuildRoutes()
    console.log('html entrypoints:')
    for (const [route, filename] of Object.entries(routes)) {
        console.log(`- "${route}" => ${filename}`)
    }
    console.log()

    const entryPoints = Object.values(routes)

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
