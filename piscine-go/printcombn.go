package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	var prev bool = false
	for i1 := '0'; i1 <= '9'; i1++ {
		n--
		if n > 0 {
			for i2 := i1 + 1; i2 <= '9'; i2++ {
				n--
				if n > 0 {
					for i3 := i2 + 1; i3 <= '9'; i3++ {
						n--
						if n > 0 {
							for i4 := i3 + 1; i4 <= '9'; i4++ {
								n--
								if n > 0 {
									for i5 := i4 + 1; i5 <= '9'; i5++ {
										n--
										if n > 0 {
											for i6 := i5 + 1; i6 <= '9'; i6++ {
												n--
												if n > 0 {
													for i7 := i6 + 1; i7 <= '9'; i7++ {
														n--
														if n > 0 {
															for i8 := i7 + 1; i8 <= '9'; i8++ {
																n--
																if n > 0 {
																	for i9 := i8 + 1; i9 <= '9'; i9++ {
																		// if n > 0 {
																		// 	for i9 := i8 + 1; i9 <= '9'; i9++ {
																		// 	}
																		// } else {
																		if i1 < i2 && i2 < i3 && i3 < i4 && i4 < i5 && i5 < i6 && i6 < i7 && i7 < i8 && i8 < i9 {
																			printCommas(prev)
																			z01.PrintRune(i1)
																			z01.PrintRune(i2)
																			z01.PrintRune(i3)
																			z01.PrintRune(i4)
																			z01.PrintRune(i5)
																			z01.PrintRune(i6)
																			z01.PrintRune(i7)
																			z01.PrintRune(i8)
																			z01.PrintRune(i9)
																			prev = true
																		}
																		// }
																		// n++
																	}
																} else {
																	if i1 < i2 && i2 < i3 && i3 < i4 && i4 < i5 && i5 < i6 && i6 < i7 && i7 < i8 {
																		printCommas(prev)
																		z01.PrintRune(i1)
																		z01.PrintRune(i2)
																		z01.PrintRune(i3)
																		z01.PrintRune(i4)
																		z01.PrintRune(i5)
																		z01.PrintRune(i6)
																		z01.PrintRune(i7)
																		z01.PrintRune(i8)
																		prev = true
																	}
																}
																n++
															}
														} else {
															if i1 < i2 && i2 < i3 && i3 < i4 && i4 < i5 && i5 < i6 && i6 < i7 {
																printCommas(prev)
																z01.PrintRune(i1)
																z01.PrintRune(i2)
																z01.PrintRune(i3)
																z01.PrintRune(i4)
																z01.PrintRune(i5)
																z01.PrintRune(i6)
																z01.PrintRune(i7)
																prev = true
															}
														}
														n++
													}
												} else {
													if i1 < i2 && i2 < i3 && i3 < i4 && i4 < i5 && i5 < i6 {
														printCommas(prev)
														z01.PrintRune(i1)
														z01.PrintRune(i2)
														z01.PrintRune(i3)
														z01.PrintRune(i4)
														z01.PrintRune(i5)
														z01.PrintRune(i6)
														prev = true
													}
												}
												n++
											}
										} else {
											if i1 < i2 && i2 < i3 && i3 < i4 && i4 < i5 {
												printCommas(prev)
												z01.PrintRune(i1)
												z01.PrintRune(i2)
												z01.PrintRune(i3)
												z01.PrintRune(i4)
												z01.PrintRune(i5)
												prev = true
											}
										}
										n++
									}
								} else {
									if i1 < i2 && i2 < i3 && i3 < i4 {
										printCommas(prev)
										z01.PrintRune(i1)
										z01.PrintRune(i2)
										z01.PrintRune(i3)
										z01.PrintRune(i4)
										prev = true
									}
								}
								n++
							}
						} else {
							if i1 < i2 && i2 < i3 {
								printCommas(prev)
								z01.PrintRune(i1)
								z01.PrintRune(i2)
								z01.PrintRune(i3)
								prev = true
							}
						}
						n++
					}
				} else {
					if i1 < i2 {
						printCommas(prev)
						z01.PrintRune(i1)
						z01.PrintRune(i2)
						prev = true
					}
				}
				n++
			}
		} else {
			printCommas(prev)
			z01.PrintRune(i1)
			prev = true
		}
		n++
	}
	z01.PrintRune('\n')
}

func printCommas(prev bool) {
	if prev == true {
		z01.PrintRune(',')
		z01.PrintRune(' ')
		prev = false
	}
}
