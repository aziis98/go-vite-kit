import { join } from 'path'
import express from 'express'
import { createServer as createViteServer } from 'vite'
import { readFile } from 'fs/promises'

import { getDevelopmentRoutes } from './routes.js'

async function createServer() {
    const app = express()

    // In middleware mode, if you want to use Vite's own HTML serving logic
    // use `'html'` as the `middlewareMode` (ref https://vitejs.dev/config/#server-middlewaremode)
    const vite = await createViteServer({
        server: { middlewareMode: 'html' },
    })

    const routes = await getDevelopmentRoutes()

    console.log(`Mounting static routes:`)
    for (const [route, file] of Object.entries(routes)) {
        const filePath = join('./frontend', file)
        console.log(`- "%s" => %s`, route, filePath)

        app.get(route, async (req, res) => {
            const htmlFile = await readFile(filePath, 'utf8')
            const htmlViteHooksFile = await vite.transformIndexHtml(req.originalUrl, htmlFile)

            res.setHeader('Content-Type', 'text/html')
            return res.send(htmlViteHooksFile)
        })
    }

    app.use(vite.middlewares)

    console.log('Started dev server on port :3000')
    app.listen(3000)
}

createServer()
