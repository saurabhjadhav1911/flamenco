/* ***** BEGIN GPL LICENSE BLOCK *****
 *
 * Original Code Copyright (C) 2022 Blender Foundation.
 *
 * This file is part of Flamenco.
 *
 * Flamenco is free software: you can redistribute it and/or modify it under
 * the terms of the GNU General Public License as published by the Free Software
 * Foundation, either version 3 of the License, or (at your option) any later
 * version.
 *
 * Flamenco is distributed in the hope that it will be useful, but WITHOUT ANY
 * WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
 * A PARTICULAR PURPOSE.  See the GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along with
 * Flamenco.  If not, see <https://www.gnu.org/licenses/>.
 *
 * ***** END GPL LICENSE BLOCK ***** */

print("Blender Render job submitted");
print("job: ", job);

const { created, settings } = job;

// Set of scene.render.image_settings.file_format values that produce
// files which FFmpeg is known not to handle as input.
const ffmpegIncompatibleImageFormats = new Set([
    "EXR",
    "MULTILAYER", // Old CLI-style format indicators
    "OPEN_EXR",
    "OPEN_EXR_MULTILAYER", // DNA values for these formats.
]);

// The render path contains a filename pattern, most likely '######' or
// something similar. This has to be removed, so that we end up with
// the directory that will contain the frames.
const renderOutput = path.dirname(settings.render_output);
const finalDir = path.dirname(renderOutput);
const renderDir = intermediatePath(finalDir);

// Determine the intermediate render output path.
function intermediatePath(render_path) {
    const basename = path.basename(render_path);
    const name = `${basename}__intermediate-${created}`;
    return path.join(path.dirname(render_path), name);
}

function frameChunker(frames, callback) {
    // TODO: actually implement.
    callback("1-10");
    callback("11-20");
    callback("21-30");
}

function authorRenderTasks() {
    let renderTasks = [];
    frameChunker(settings.frames, function(chunk) {
        const task = author.Task(`render-${chunk}`);
        const command = author.Command("blender-render", {
            cmd: settings.blender_cmd,
            filepath: settings.filepath,
            format: settings.format,
            render_output: path.join(renderDir, path.basename(renderOutput)),
            frames: chunk,
        });
        task.addCommand(command);
        renderTasks.push(task);
    });
    return renderTasks;
}

function authorCreateVideoTask() {
    if (ffmpegIncompatibleImageFormats.has(settings.format)) {
        return;
    }
    if (!settings.fps || !settings.output_file_extension) {
        return;
    }

    const stem = path.stem(settings.filepath).replace('.flamenco', '');
    const outfile = path.join(renderDir, `${stem}-${settings.frames}.mp4`);

    const task = author.Task('create-video');
    const command = author.Command("create-video", {
        input_files: path.join(renderDir, `*${settings.output_file_extension}`),
        output_file: outfile,
        fps: settings.fps,
    });
    task.addCommand(command);

    print(`Creating output video for ${settings.format}`);
    return task;
}

const renderTasks = authorRenderTasks();
const videoTask = authorCreateVideoTask(renderTasks);

if (videoTask) {
    // If there is a video task, all other tasks have to be done first.
    for (const rt of renderTasks) {
        videoTask.addDependency(rt);
    }
    job.addTask(videoTask);
}
for (const rt of renderTasks) {
    job.addTask(rt);
}