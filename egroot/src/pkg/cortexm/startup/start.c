#include "runtime/types.h"

extern uint32 __dataStart, __dataEnd, __dataStartFlash;
extern uint32 __bssStart, __bssEnd;

void cortexm_startup_setupDataBSS() {
	uint32 *src = &__dataStartFlash;
	uint32 *dst = &__dataStart;
	uint32 *end = &__dataEnd;

	while (dst < end) {
		*dst = *src;
		++dst;
		++src;
	}

	dst = (uint32 *) & __bssStart;
	end = (uint32 *) & __bssEnd;

	while (dst < end) {
		*dst = 0;
		++dst;
	}
}

static void nmi_handler(void) {
	for (;;);
}

static void hardfault_handler(void) {
	for (;;);
}

extern uint32 stackptr;
int main();

uint32 *cortexm_startup_vectors[4] __attribute__ ((section("vectors"))) = {
	&stackptr,
	(uint32 *) main,              // code entry point
	(uint32 *) nmi_handler,       // NMI handler (not really)
	(uint32 *) hardfault_handler, // hard fault handler 
};