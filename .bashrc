# Software replacement for the 'beep' command
function beep() {
    local freq=440
    local len=0.2
    local reps=1
    local delay=0.1

    OPTIND=1
    while getopts "f:l:r:d:" opt; do
        case $opt in
            f) freq=$OPTARG ;;
            l) len=$(awk "BEGIN {print $OPTARG / 1000}") ;; # Convert ms to seconds
            r) reps=$OPTARG ;;
            d) delay=$(awk "BEGIN {print $OPTARG / 1000}") ;; # Convert ms to seconds
        esac
    done

    for ((i=0; i<reps; i++)); do
        # Generate a sine wave tone silently (-q)
        play -q -n synth "$len" sine "$freq"

        # Add delay between beeps, but not after the last one
        if [ $((i + 1)) -lt "$reps" ]; then
            sleep "$delay"
        fi
    done
}
