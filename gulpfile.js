var gulp        = require('gulp');
var gutil       = require('gulp-util');

var less        = require('gulp-less');
var include     = require('gulp-include')
var notify      = require("gulp-notify");
var rename      = require('gulp-rename');
var uglify      = require('gulp-uglify');
var uglifycss   = require('gulp-uglifycss');

// PATHS

var inputDir        = '_src';
var outputDir       = '_ass/v1';

var cssInputDir     = inputDir  + '/less/*main.less';
var cssWatchDir     = inputDir  + '/less/*.less';
var cssOutputDir    = outputDir + '/css';

// FUNCTIONS

function processCSS(prod) {
    return gulp.src(cssInputDir)
        .pipe(less().on('error', processErrors))
        .pipe(prod ? uglifycss() : gutil.noop())
        .pipe(gulp.dest(cssOutputDir))
    ;
}

function processErrors() {
    var args = Array.prototype.slice.call(arguments);

    notify
        .onError({
            title: "Compile Error",
            message: "<%= error.message %>"
        })
        .apply(this, args)
    ;

    this.emit('end');
}


// DEFAULT

gulp.task('default', ['watch']);

gulp.task('watch', function() {
    gulp.watch(cssWatchDir, ['css']);
});

gulp.task('css', function () {
    return processCSS(false);
});
