const ws = require('windows-shortcuts');
const asar = require("@electron/asar");
const fs = require("fs");
const path = require("path");
const os = require('os');


function patchTestwe(p) {
    let skip = false;
    let patched = [];

    fs.mkdtemp(path.join(os.tmpdir(), 'testwe'), (err, folder) => {
        if (err) throw err;
        asar.extractAll(path.join(p, "resources", "app.asar"), path.join(folder, "app"));
        const code = fs.readFileSync(path.join(folder, "app", "dist", "electron", "main.js"), { encoding: 'utf8' }).split(/\r?\n/);

        for (line of code) {
            if (!skip) {
                patched.push(line);
            }

            if (line.startsWith("async function detectVirtualEnvironemnt")) {
                skip = true;
            }

            if (skip && line == "}") {
                patched.push("}");
                skip = false;
            }
        }

        fs.writeFileSync(path.join(folder, "app", "dist", "electron", "main.js"), patched.join("\n"));
        asar.createPackage(path.join(folder, "app"), path.join(p, "resources", "app.asar")).then(() => console.log("Patching is done !"));
    })
}


if (require.main == module) {
    ws.query(path.join(process.env.USERPROFILE, "Desktop", "Testwe.lnk"), function (err, opts) {
        if (err !== null || opts.workingDir === null) {
            console.log("Couldn't find shortcut on Desktop named `TestWe`");
            console.log("The easiest way to fix this: Reinstall TestWe")
            process.exit(1);
        }

        patchTestwe(opts.workingDir);
    });
}