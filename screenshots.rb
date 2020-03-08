
`ls fixtures`.split("\n").each do |fixture|
  `gggv &`
  sleep 1
  `go run gggv-vhs if #{fixture} ou :4200`
  sleep 1
  `screencapture -l $(GetWindowId gggv gggv) screenshots/#{fixture.gsub('.json', '.png')}`
  `kill -9 $(pidof gggv)`
end
