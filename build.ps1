# build.ps1 — Markit build script

Write-Host "Building Markit..." -ForegroundColor Cyan

# Step 1 — build exe
& "C:\Users\Bilal\go\bin\wails.exe" build

# Step 2 — create installer
& "C:\Program Files (x86)\NSIS\makensis.exe" /DARG_WAILS_AMD64_BINARY="D:\golearn\projects\md-file-reader\build\bin\Markit.exe" "D:\golearn\projects\md-file-reader\build\windows\installer\project.nsi"

Write-Host "Done! Check build/bin/" -ForegroundColor Green