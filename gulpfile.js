const gulp = require("gulp");
const sass = require("gulp-sass")(require("sass"));
const cleanCSS = require("gulp-clean-css");
const rename = require("gulp-rename");
const terser = require("gulp-terser");
const sourcemaps = require("gulp-sourcemaps");
const fs = require("fs");
const path = require("path");
const concat = require("gulp-concat");

/* -------------------------
   Task SCSS → CSS
------------------------- */
function stylesFrontend() {
  return gulp.src("resources/scss/frontend/*.scss")
    .pipe(sourcemaps.init())
    .pipe(sass().on("error", sass.logError))
    .pipe(gulp.dest("public/css/frontend"))          // versão normal
    .pipe(cleanCSS())                        // minificar
    .pipe(rename({ suffix: ".min" }))
    .pipe(sourcemaps.write("../maps"))       // source maps em /css/maps
    .pipe(gulp.dest("public/css/frontend/minify"));  // versão minificada
}

function stylesBackend() {
  return gulp.src("resources/scss/backend/*.scss")
    .pipe(sourcemaps.init())
    .pipe(sass().on("error", sass.logError))
    .pipe(gulp.dest("public/css/backend"))          // versão normal
    .pipe(cleanCSS())                        // minificar
    .pipe(rename({ suffix: ".min" }))
    .pipe(sourcemaps.write("../maps"))       // source maps em /css/maps
    .pipe(gulp.dest("public/css/backend/minify"));  // versão minificada
}

/* -------------------------
   Task JS → Bundles
------------------------- */
function bundleJS(done) {
  const jsDir = "resources/js";
  const bundleDir = "public/js";

  // Lê subpastas dentro de /js
  fs.readdirSync(jsDir, { withFileTypes: true })
    .filter(dirent => dirent.isDirectory())
    .forEach(dirent => {
      const subfolder = dirent.name;
      const subfolderPath = path.join(jsDir, subfolder);

      // Cria pasta de output se não existir
      const outFolder = path.join(bundleDir, subfolder);
      if (!fs.existsSync(outFolder)) fs.mkdirSync(outFolder, { recursive: true });

      // Pega todos os JS da subpasta
      gulp.src(`${subfolderPath}/*.js`)
        .pipe(concat(`bundle-${subfolder}.js`))      // cria bundle
        .pipe(gulp.dest(outFolder))                  // salva bundle normal
        .pipe(terser())                              // minifica
        .pipe(rename({ suffix: ".min" }))
        .pipe(gulp.dest(outFolder));                 // salva minificado
    });

  done();
}

/* -------------------------
   Watch SCSS e JS
------------------------- */
function watchFiles() {
  gulp.watch("resources/scss/**/*.scss", stylesFrontend);
  gulp.watch("resources/js/**/*.js", bundleJS);
}

/* -------------------------
   Export tasks
------------------------- */
exports['build-frontend'] = stylesFrontend;               // só compila SCSS
exports['build-backend'] = stylesBackend;               // só compila SCSS
exports['build-js'] = bundleJS;              // só cria JS bundles
exports['build-watch'] = gulp.series(stylesFrontend, bundleJS, watchFiles); // compila tudo + watch
