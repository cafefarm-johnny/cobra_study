package cmd

import (
	"flag"
	"fmt"
	"time"

	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/jedib0t/go-pretty/v6/progress"

	"github.com/spf13/cobra"
)

var (
	autoStop    = flag.Bool("auto-stop", true, "Auto-stop rendering?")
	numTrackers = flag.Int("num-trackers", 5, "Number of Trackers")
)

var progressCmd = &cobra.Command{
	Use:     "progress",
	Short:   "progress bar test command",
	Long:    "progress bar test command",
	Example: "progress --help",
	Run: func(cmd *cobra.Command, args []string) {
		progressAction()
	},
}

func init() {
	rootCmd.AddCommand(progressCmd)
}

func progressAction() {
	fmt.Println()
	// 프로그레스 writer 인스턴스 생성
	pw := progress.NewWriter()
	setOptions(pw)
	setStyle(pw)

	// 비동기 모드로 렌더링 호출
	go pw.Render()

	// 예제를 위해 tracker 생성
	for trackerNo := int64(1); trackerNo <= int64(*numTrackers); trackerNo++ {
		// 프로그레스 로직 실행 (멀티 다운로드 기능처럼 동시성으로 실행)
		go tracking(pw, trackerNo)

		// 각각의 tracker 실행 term (AutoStop 활성화 시 무시됨)
		if !*autoStop {
			time.Sleep(time.Second)
		}
	}

	// 하나 이상의 tracker가 활성화 될 때 까지 1초간 기다린다
	time.Sleep(time.Second)

	// 프로그레스바 렌더링
	for pw.IsRenderInProgress() {
		// 활성 tracker가 없는 경우 프로그레스 중지
		if !*autoStop && pw.LengthActive() == 0 {
			pw.Stop()
		}

		// * 모든 tracker 처리 완료 후 펜딩 시간 (AutoStop 활성화 시 무시됨)
		time.Sleep(time.Second)
	}

	fmt.Println("\nAll Process Complete!!")
}

func setOptions(pw progress.Writer) {
	pw.SetAutoStop(*autoStop)
	pw.SetTrackerLength(50)     // 프로그레스 바 width 길이
	pw.ShowOverallTracker(true) // 전체 프로세스 처리 진행률 프로그레스 바 표현 여부
	pw.ShowTime(true)           // 처리 소모 시간 표현 여부
	pw.ShowValue(true)          // unit 값 표현 여부
	pw.ShowTracker(true)        // 프로그레스 바 표현 여부
	pw.SetMessageWidth(35)
	pw.SetNumTrackersExpected(*numTrackers) // 표현할 tracker 수량 설정 (더 정확한 처리 진행률 계산이 된다고 함)
	pw.SetSortBy(progress.SortByPercentDsc)
	pw.SetStyle(progress.StyleDefault)
	pw.SetTrackerPosition(progress.PositionRight) // 프로그레스 바 우측 정렬
	pw.SetUpdateFrequency(time.Millisecond * 100) // 진행률 상태 갱신 시간
}

func setStyle(pw progress.Writer) {
	pw.Style().Colors.Percent = text.Colors{text.FgRed}
	pw.Style().Colors.Time = text.Colors{text.FgWhite}
	pw.Style().Colors.Value = text.Colors{text.FgGreen}

	pw.Style().Options.SnipIndicator = "..." // 메세지 width 가 상황에 따라 잘리는 경우 대체 표현 문자
	pw.Style().Options.Separator = " --- "   // 메세지와 프로그레스 바 사이의 구분자
	pw.Style().Options.DoneString = "100%"
	pw.Style().Options.PercentFormat = "%4.1f%%" // 진행률 포맷
}

func tracking(pw progress.Writer, trackerNo int64) {
	// 전체 프로세스 처리 진행률 주기
	totalProcessPerCycle := trackerNo * trackerNo * trackerNo * 1000 // 5 * 5 * 5 * 1000 = 125000
	// 각각 프로세스 처리 진행률 주기
	processPerCycle := trackerNo * int64(*numTrackers) * 100 // 5 * 5 * 100 = 2500

	// unit: 용량 표현 방식 (UnitsBytes: kilo byte)
	units := &progress.UnitsBytes
	message := fmt.Sprintf("Downloading \"filename-%d.ext\"", trackerNo)
	tracker := progress.Tracker{Message: message, Total: totalProcessPerCycle, Units: *units}

	pw.AppendTracker(&tracker)

	// 0.1초마다 진행률 증가 이벤트 발생
	c := time.Tick(time.Millisecond * 100)
	for !tracker.IsDone() {
		select {
		case <-c:
			tracker.Increment(processPerCycle)
		}
	}
}
