import { dirname, resolve } from 'path'
import express from 'express'
import { createServer as createViteServer } from 'vite'
import { fileURLToPath } from 'url'

const __dirname = dirname(fileURLToPath(import.meta.url))

async function createServer(customHtmlRoutes) {
    const app = express()

    // In middleware mode, if you want to use Vite's own HTML serving logic
    // use `'html'` as the `middlewareMode` (ref https://vitejs.dev/config/#server-middlewaremode)
    const vite = await createViteServer({
        server: { middlewareMode: 'html' },
    })

    for (const [route, file] of Object.entries(customHtmlRoutes)) {
        app.get(route, (req, res) => {
            const filePath = resolve(__dirname, file)
            console.log(`Custom Route: %s`, req.url)

            return res.sendFile(filePath)
        })
    }

    app.use(vite.middlewares)

    app.listen(3000)
}

createServer({
    // Useful for developing with dynamic routes
    '/': './index.html',
})
