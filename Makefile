day = $(shell date +'%-d')

new:
	@echo "Creating new file structure for day" $(day)"..."

	@if [ $(day) -lt 10 ] ; then \
  		mkdir calendar/day-0$(day); \
  		cp template calendar/day-0$(day)/part1.go; \
  		cp template calendar/day-0$(day)/part2.go; \
  	else \
  		mkdir calendar/day-$(day); \
		cp template calendar/day-$(day)/part1.go; \
		cp template calendar/day-$(day)/part2.go; \
    fi

	@git add calendar/

