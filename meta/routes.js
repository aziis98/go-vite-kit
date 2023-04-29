import { spawn } from 'child_process'

function transformRoutes(entrypoints) {
    return Object.fromEntries(entrypoints.map(({ route, filename }) => [route, filename]))
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

    console.log('loading build entrypoints...')
    return transformRoutes(await readCommandOutputAsJSON('go run ./meta/routes'))
}
