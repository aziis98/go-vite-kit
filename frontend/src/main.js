async function main() {
    const res = await fetch('/api/status')

    if (res.ok) {
        console.log('Server online')
    }
}

main()
