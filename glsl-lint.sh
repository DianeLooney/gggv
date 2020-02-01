wget --no-verbose https://github.com/angular/clang-format/raw/master/bin/linux_x64/clang-format
chmod a+x clang-format
./clang-format --version
export IFS=$'\n'
for file in $(git ls-files | grep \\.glsl$); do
    echo Formatting "$file"
    ./clang-format -style=Chromium -i "$file"
done
echo -e "\\n\\n\\n\\tChecking diff...\\n\\n\\n"
set -e
git diff --exit-code
echo -e "\\tStyle is fine"
