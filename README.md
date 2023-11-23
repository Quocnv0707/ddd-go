Cấu trúc trong một ứng dụng Go:
golang_project_structure/
    api/
     |
    assets/
     |
    build/
     |________ci/
     |
     |________package/
     |
    cmd/
     |_________your_app/
     |          |
     |          |_____main.go
     |
    configs/
     |
    deployments/
     |
    docs/
     |
    examples/
     |
    githooks/
     |
    init/
     |
    internal/
     |   
     |_____app/your_app/
     |          |
     |          |_____pkg/your_private_lib
     |
    pkg/
     |
     |_____your_package_lib/
     |
    scripts/
     |
    test/
     |
    third-party/
     |
    tools/
    |
    vendor/
     |
    web/
     |_____app/
     |_____static/
     |_____template/
     |
     |
    website/

Giải thích: 
/cmd
    Thư mục chứa các ứng dụng chính cho dự án này.

    Tên thư mục cho mỗi ứng dụng phải khớp với tên của tập tin thực thi mà bạn muốn có (ví dụ: /cmd/myapp).

/internal
    Gói này chứa mã thư viện riêng được sử dụng trong dịch vụ của bạn, nó dành riêng cho chức năng của dịch vụ và không được chia sẻ với các dịch vụ khác. Ngoài ra, bạn vẫn có thể tách mã nội bộ được chia sẻ và không được chia sẻ.
    /internal/app/myapp/: Đây là nơi lưu trữ mã nguồn cho từng service cụ thể, bao gồm các logic xử lý và chức năng.
    /pkg/your_package_lib: nơi chứa các thư viện, mã nguồn không muốn chia sẻ với người khác

/pkg
    Thư mục này chứa mã phù hợp cho các dịch vụ khác sử dụng, mã này có thể bao gồm các ứng dụng khách API hoặc các chức năng tiện ích có thể hữu ích cho các dự án khác nhưng không biện minh cho dự án của riêng họ.

/vendor
    Thư mục chứa các phụ thuộc của ứng dụng (được quản lý thủ công hoặc bằng công cụ quản lý phụ thuộc ưa thích của bạn tương tự như tính năng tích hợp mới là Go Modules). Câu lệnh go mod vendor sẽ tạo ra cho bạn một thư mục /vendor. Lưu ý rằng bạn có thể sẽ cần phải thêm cờ hiệu -mod=vendor cho câu lệnh go build nếu bạn không dùng Go 1.14, phiên bản được thêm cờ hiệu mặc định.

    Không nên commit các phụ thuộc ứng dụng nếu bạn đang muốn xây dựng một thư viện.

    Lưu ý rằng kể từ phiên bản 1.13, Go bật tính năng module proxy (mặc định dùng máy chủ module proxy https://proxy.golang.org). Đọc thêm để xem liệu nó có phù hợp với tất cả các yêu cầu và ràng buộc của bạn hay không ở đây. Nếu có thì bạn không cần tới thư mục vendor.

/api
    Thư mục chứa bản mô tả OpenAPI/Swagger, tập tin lược đồ JSON, tập tin định nghĩa giao thức.

Thư mục ứng dụng Web
/web
    Thư mục chứa các thành phần cụ thể của ứng dụng web: tài nguyên web tĩnh, mẫu bên máy chủ và SPAs.

    Để các tập mẫu confd và consul-template ở đây.

/init
    Thư mục chứa phần khởi tạo hệ thống (systemd, upstart, sysv) và cấu hình quản lý/giám sát tiến trình (runit, supervisord).

/scripts
    Thư mục chứa tập lệnh để thực hiện các hoạt động xây dựng, cài đặt, phân tích ...

    Các tập lệnh này làm cho tập Makefile ở cấp cao nhất nhỏ gọn và đơn giản. 
    (Ví dụ: https://github.com/hashicorp/terraform/blob/master/Makefile)

/build
    Thư mục chứa gói và tích hợp liên tục.

    Đặt các cấu hình và tập lệnh các gói đám mây (AMI), container (Docker), OS (deb, rpm, pkg) của bạn vào thư mục /build/package.

    Đặt cấu hình và tập lệnh CI (travis, circle, drone) trong thư mục /build/ci. Lưu ý rằng một vài công cụ CI (ví dụ: Travis CI) rất kén chọn vị trí của tập tin cấu hình. Thử đặt các tập tin cấu hình ở thư mục /build/ci, lên kết chúng với vị trí mà công cụ CI mong đợi (khi có thể).

/deployments
    Thư mục chứa IaaS, PaaS, các cấu hình và mẫu triển khai điều phối hệ thống và vùng chứa (docker-compose, kubernetes/helm, mesos, terraform, bosh). 
    Lưu ý rằng trong một số repo (đặc biệt là các ứng dụng được triển khai với kubernetes) thư mục này được gọi là /deploy.

/test
    Thư mục chứa các ứng dụng thử nghiệm bên ngoài bổ sung và dữ liệu thử nghiệm. Hãy thoải mái cấu trúc thư mục /test theo cách bạn muốn. Đối với các dự án lớn hơn, điều hợp lý là có một thư mục con dữ liệu. Ví dụ: bạn có thể có /test/data hoặc /test/testdata nếu bạn cần Go bỏ qua những gì trong thư mục đó. Lưu ý rằng Go cũng sẽ bỏ qua các thư mục hoặc tệp bắt đầu bằng "." hoặc "_", vì vậy bạn có thể linh hoạt hơn về cách đặt tên cho thư mục dữ liệu thử nghiệm của mình.

Thư mục khác
/docs
    Thư mục chứa tài liệu người dùng và bản thiết kế (bên cạnh tài liệu do godoc tạo ra).

/tools
    Thư mục chứa công cụ hỗ trợ cho dự án này. Lưu ý rằng các công cụ này có thể nhập mã từ thư mục /pkg và /internal.

/examples
    Thư mục chứa mẫu cho ứng dụng và/hoặc các thư viện công cộng của bạn.

/third_party
    Thư mục chứa các công cụ trợ giúp bên ngoài, mã phân nhánh và các tiện ích bên thứ ba khác (ví dụ: giao diện người dùng Swagger).

/githooks
    Thư mục chứa git hooks.

/assets
    Các tài sản khác đi cùng với kho lưu trữ của bạn (hình ảnh, logo ...).

/website
    Đây là nơi để dữ liệu trang web của bạn nếu bạn không sử dụng các trang của GitHub.


Cấu trúc của một dự án microservice
api/: Có thể chứa các file hoặc package định nghĩa các API endpoints hoặc giao diện người dùng.

cmd/your_app/: Đây có thể là nơi triển khai các service hoặc ứng dụng cụ thể. Mỗi thư mục trong cmd có thể tương ứng với một microservice.

configs/: Lưu trữ các tập tin cấu hình cho từng service hoặc ứng dụng.

internal/app/your_app/: Đây có thể là nơi lưu trữ mã nguồn cho từng service cụ thể, bao gồm các logic xử lý và chức năng.

pkg/: Chứa các package chung có thể được chia sẻ giữa các service. Điều này giúp tạo ra các thư viện tái sử dụng và giảm sự phụ thuộc giữa các service.

test/: Chứa các test unit hoặc test tích hợp cho từng service.

vendor/: Có thể chứa dependencies hoặc các thư viện cần thiết cho dự án.

web/: Có thể chứa các thành phần liên quan đến giao diện người dùng hoặc phần frontend của ứng dụng.

===================================
- Phân tách chức năng: Cả 2 cấu trúc golang và microservice đều tập trung vào việc phân chia chức năng thành các phần nhỏ, dễ quản lý và mở rộng.
- Tổ chức logic theo service: cả hai đều đẩy việc tổ chức logic và chức năng thành các service độc lập
- Module hoá và tái sử dụng: cả hai đều khuyến khích tạo ra các modle có thể tái sử dụng, giúp cải thiện tính tái sử dụng và quản lý codebase.