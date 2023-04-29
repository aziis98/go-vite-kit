import { spawn } from 'child_process'
import axios from 'axios'

function transformRoutes(entrypoints) {
    return Object.fromEntries(entrypoints.map(({ route, filename }) => [route, filename]))
}

export async function getDevelopmentRoutes() {
    const res = await axios.get('http://127.0.0.1:4000/dev/routes')
    return transformRoutes(res.data)
}

export async function getBuildRoutes() {
    // Thanks to ChatGPT
    function readCommandOutputAsJSON(command) {
        const [cmd, ...args] = command.split(' ')

        return new Promise((resolve, reject) => {
            const child = spawn(cmd, args)

            let stdout = ''

            child.stdout.on('data', data => {
                stdout += data.toString()
            })

            child.on('close', code => {
                if (code !== 0) {
                    reject(`Command ${cmd} ${args.join(' ')} failed with code ${code}`)
                    return
                }

                try {
                    const output = JSON.parse(stdout)
                    resolve(output)
                } catch (e) {
                    reject(`Error parsing JSON output: ${e.message}`)
                }
            })
        })
    }

    return transformRoutes(await readCommandOutputAsJSON('go run ./cmd/routes'))
}
